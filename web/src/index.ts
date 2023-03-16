// Import the necessary modules
import { Request, Response } from "express";

// Define a function to handle HTTP requests
export function handleRequest(req: Request, res: Response): void {
  // Display a message in the browser console
  console.log("Hello, world!");

  // Display a message on the webpage
  const message = document.createElement("div");
  message.textContent = "Hello, world!";
  document.body.appendChild(message);
}
