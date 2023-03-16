"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.handleRequest = void 0;

function handleRequest(req, res) {
    console.log("Hello, world!");
    var message = document.createElement("div");
    message.textContent = "Hello, world!";
    document.body.appendChild(message);
}
exports.handleRequest = handleRequest;
