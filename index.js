const express = require("express");

const app = express();

app.get("/", (req, res) => {
  res.status(200).json({ status: "success", message: "Hello, Codespaces!" });
});

app.listen(8080, () => {
  console.log("App is listening on port 8080.");
});
