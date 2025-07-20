import { Octokit } from "octokit"

const owner = 'phatelot';
const repo = 'fehwgc-locator';
const branch = 'main';

let octokit: Octokit;
let token: string = (() => {
    const token = getGithubToken() || '';
    if (token) {
      octokit = new Octokit({
        auth: token,
      });
    }
    return token;
})();

export function octokitReady(): boolean {
	return !!octokit;
}

export function updateGithubToken(newToken: string) {
	token = newToken;
	localStorage.setItem('fehwgc-locator-admin', token);
	octokit = new Octokit({
		auth: token,
	});
}

export function getGithubToken(): string | null {
	return localStorage.getItem('fehwgc-locator-admin');
}

export async function triggerUpdateImageWorkflow() {
	if (!octokitReady()) {
		throw new Error('please fill in Github token')
	}

	return await octokit.request('POST /repos/{owner}/{repo}/actions/workflows/{workflow_id}/dispatches', {
		owner,
		repo,
		workflow_id: 'downloadImage.yml',
		ref: branch,
		headers: {
		  'X-GitHub-Api-Version': '2022-11-28'
		}
	  })
}
