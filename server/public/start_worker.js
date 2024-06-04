if ("serviceWorker" in navigator) {
	navigator.serviceWorker.register("resources/sw.js")
		.then((reg) => {
			reg.addEventListener("statechange", (event) => {
				console.log("received `statechange` event", { reg, event });
			});
			console.log("service worker registered", reg);
			reg.active.postMessage({ type: "clientattached" });
		}).catch((err) => {
			console.error("service worker registration failed", err);
		});
	navigator.serviceWorker.addEventListener("controllerchange", (event) => {
		console.log("received `controllerchange` event", event);
	});
} else {
	console.error(
		"serviceWorker is missing from `navigator`. Note service workers must be served over https or on localhost",
	);
}
