# gen-metafiles

Generate Metafiles (gen-metafiles) is a basic example of an agent-based GitHub Copilot Extension. It responds to simple identifiers in prompts, using Copilot's LLM API, and creates either examples of .gitignore and .gitattributes file or example model files with the user's name and year filled in (if appropriate). This repository can serve as an example of the building blocks of a Copilot Extension. See [index.js](/index.js) for the main logic. (It is originally based off of the Blackbeard extension demo code.)

> [!NOTE]
> Copilot Extensions are in public preview and may be subject to change.
> 
> All enrolled users with a GitHub Copilot Individual subscription can use Copilot Extensions.
> For organizations and enterprises with a Copilot Business or Copilot Enterprise subscription, organization owners and enterprise administrators can grant access to Copilot Extensions for use within their company.

## Development

Install dependencies:

```bash
npm install
```

To run:

```bash
npm start
```

or in watch mode:

```bash
npm run dev
```

## Documentation
- [Using Copilot Extensions](https://docs.github.com/en/copilot/using-github-copilot/using-extensions-to-integrate-external-tools-with-copilot-chat)
- [About building Copilot Extensions](https://docs.github.com/en/copilot/building-copilot-extensions/about-building-copilot-extensions)
- [Set up process](https://docs.github.com/en/copilot/building-copilot-extensions/setting-up-copilot-extensions)
- [Communicating with the Copilot platform](https://docs.github.com/en/copilot/building-copilot-extensions/building-a-copilot-agent-for-your-copilot-extension/configuring-your-copilot-agent-to-communicate-with-the-copilot-platform)
- [Communicating with GitHub](https://docs.github.com/en/copilot/building-copilot-extensions/building-a-copilot-agent-for-your-copilot-extension/configuring-your-copilot-agent-to-communicate-with-github)

