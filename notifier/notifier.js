const path = require("path");
const express = require("express");
const app = express();
const notifier = require("node-notifier");
const port = process.env.PORT || 9000;

app.use(express.json());

app.get("/health", (req, res) => {
  return res.status(200).send("ok");
});

app.post("/notify", (req, res) => {
  notify(req.body, (reply) => res.send(reply));
  // return res.status(200).send("ok");
});

app.listen(port, () => {
  console.log(`Server listening on port ${port}`);
});

const notify = ({ title, message }, cb) => {
  notifier.notify(
    {
      title: title || "Unknown title",
      message: message || "Unknown message",
      icon: path.join(__dirname, "icon.png"),
      sound: true,
      wait: true,
      reply: true,
      closeLabel: "Completed?",
      timeout: 15,
    },
    (err, response, reply) => {
      cb(reply);
    }
  );
};
