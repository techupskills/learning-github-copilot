import * as vscode from 'vscode';

const BASE_PROMPT = 'You are a helpful API finder. Your job is to locate and provide documentation on the top 3 APIs that match the prompt. Respond with basic documentation about the API, how to call it, and its parameters. Cite your sources and provide links to more information where appropriate.  Use bullet points in explanations wherever possible. If the user asks a question that is not relevant to APIs, politely decline to respond.';

const SAMPLE_CODE_PROMPT = 'You are a helpful source of API examples. Your job is to locate the top 3 APIs that match the prompt and provide simple example code for how to call them. Keep the code as short as possible to still provide a usable, relevant example. If the user asks a question that is not relevant to APIs, politely decline to respond.';

export function activate(context: vscode.ExtensionContext) {

	// define a chat handler
	const handler: vscode.ChatRequestHandler = async (request: vscode.ChatRequest, context: vscode.ChatContext, stream: vscode.ChatResponseStream, token: vscode.CancellationToken) => {

		// initialize the prompt
		let prompt = BASE_PROMPT;

		if (request.command === 'examples') {
			prompt = SAMPLE_CODE_PROMPT;
		}

		// initialize the messages array with the prompt
		const messages = [
			vscode.LanguageModelChatMessage.User(prompt),
		];

		// get all the previous participant messages
		const previousMessages = context.history.filter(
			(h) => h instanceof vscode.ChatResponseTurn
		);

		// add the previous messages to the messages array
		previousMessages.forEach((m) => {
			let fullMessage = '';
			m.response.forEach((r) => {
				const mdPart = r as vscode.ChatResponseMarkdownPart;
				fullMessage += mdPart.value.value;
			});
			messages.push(vscode.LanguageModelChatMessage.Assistant(fullMessage));
		});

		// add in the user's message
		messages.push(vscode.LanguageModelChatMessage.User(request.prompt));

		// send the request
		const chatResponse = await request.model.sendRequest(messages, {}, token);

		// stream the response
		for await (const fragment of chatResponse.text) {
			stream.markdown(fragment);
		}

		return;

	};

	// create participant
	const apifinder = vscode.chat.createChatParticipant("chat-tutorial.api-finder", handler);

	// add icon to participant
	apifinder.iconPath = vscode.Uri.joinPath(context.extensionUri, 'apifinder.jpeg');
}

export function deactivate() { }
