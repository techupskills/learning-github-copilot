import { Octokit } from "@octokit/core";
import express from "express";
import { Readable } from "node:stream";

const app = express()

app.get("/", (req, res) => {
  res.send("Welcome to the Metafiles GitHub Copilot Extension!")
});

app.post("/", express.json(), async (req, res) => {
  // Identify the user, using the GitHub API token provided in the request headers.
  const tokenForUser = req.get("X-GitHub-Token");
  const octokit = new Octokit({ auth: tokenForUser });
  const user = await octokit.request("GET /user");
  console.log("User:", user.data);



// Define the URL for the API endpoint to fetch the current user's information
const apiUrl = 'https://api.github.com/user';



  // Parse the request payload and log it.
  const payload = req.body;
  console.log("Payload:", payload);


  
  // Define the LLM's role and what it should look for and do
  const messages = payload.messages;
  messages.unshift({
    role: "system",
    content: "Search for a programming language name in the message from the user. If you find one, generate example .gitignore and .gitattributes files for that programming language",
  });
  messages.unshift({
    role: "system",
    content: "Search for an open source software license name, type, or abbreviation in the message from the user. If you find one, do the following: Provide a few key bullet points about the license and when it should/should not be used. Then generate an example LICENSE file for that license. If the example file contains the text [NAME] or [YOUR NAME], replace that text with " + user.data.name + ". If the LICENSE file contains the text [YEAR], replace [YEAR] with the current year.",
  });
 

  // Use Copilot's LLM to generate a response to the user's messages, with
  // our extra system messages attached.
  const copilotLLMResponse = await fetch(
    "https://api.githubcopilot.com/chat/completions",
    {
      method: "POST",
      headers: {
        authorization: `Bearer ${tokenForUser}`,
        "content-type": "application/json",
      },
      body: JSON.stringify({
        messages,
        stream: true,
      }),
    }
  );

  // Stream the response straight back to the user.
  Readable.from(copilotLLMResponse.body).pipe(res);
})

const port = Number(process.env.PORT || '3000')
app.listen(port, () => {
  console.log(`Server running on port ${port}`)
});
