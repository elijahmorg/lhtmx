navigator.serviceWorker.register("sw.js");

// async function hello() {
//   const name = document.querySelector("#name").value;
//
//   const res = await fetch("api/hello", {
//     method: "POST",
//     headers: {
//       "Content-Type": "application/json",
//     },
//     body: JSON.stringify({ name }),
//   });
//
//   const { message } = await res.json();
//
//   alert(message);
// }
