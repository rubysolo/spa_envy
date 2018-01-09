var valueEl = document.getElementById("value");

if (config && config.RUNTIME) {
    valueEl.innerText = " = " + config.RUNTIME;
} else {
    valueEl.innerText = " is not set."
}