import { Request, Response } from "express";

export function handleRequest(req: Request, res: Response): void {
  console.log("Hello, world!");

  const message = document.createElement("div");
  message.textContent = "Hello, world!";
  document.body.appendChild(message);
}
