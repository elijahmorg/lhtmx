importScripts("wasm_exec.js");
// importScripts(
//   "https://cdn.jsdelivr.net/gh/nlepage/go-wasm-http-server@v1.1.0/sw.js",
// );

addEventListener("install", (event) => {
  event.waitUntil(skipWaiting());
});

addEventListener("activate", (event) => {
  event.waitUntil(clients.claim());
});

registerWasmHTTPListener("main.wasm", { base: "" });

function registerWasmHTTPListener(wasm, { base, args = [] } = {}) {
  let path = new URL(registration.scope).pathname;
  if (base && base !== "") {
    path = `${trimEnd(path, "/")}/${trimStart(base, "/")}`;
  }

  const handlerPromise = new Promise((setHandler) => {
    self.wasmhttp = {
      path,
      setHandler,
    };
  });

  const go = new Go();
  go.argv = [wasm, ...args];
  WebAssembly.instantiateStreaming(fetch(wasm), go.importObject).then((
    { instance },
  ) => go.run(instance));

  addEventListener("fetch", (e) => {
    const url = new URL(e.request.url);
    const { pathname } = new URL(e.request.url);
    console.log(e);
    console.log("pathname: ", pathname);
    console.log("path: ", path);
    if (
      pathname === "/wasm_exec.js" || pathname == "/sw.js" ||
      pathname === "/start_worker.js"
    ) {
      e.respondWith(fetch(e.request));
      return;
    } else if (url.hostname === "localhost") {
      e.respondWith(handlerPromise.then((handler) => handler(e.request)));
    } else {
      // For requests to other domains, just pass them along to the network
      e.respondWith(fetch(e.request));
    }
    // if (!pathname.startsWith(path)) {
    //   console.log("fallback to network");
    //   return fetch(request);
    // }
  });
}

function trimStart(s, c) {
  let r = s;
  while (r.startsWith(c)) r = r.slice(c.length);
  return r;
}

function trimEnd(s, c) {
  let r = s;
  while (r.endsWith(c)) r = r.slice(0, -c.length);
  return r;
}
