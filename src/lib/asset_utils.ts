export function getPublicImageLink(imageFileName: string) : string {
	return normalize(import.meta.env.BASE_URL + "/images/" + imageFileName);
}

function normalize(url: string): string {
	return url.replace('//', '/')
}
