type Position = {
	scale: number;
	x: number;
	y: number;
}

export const positionsByCharacterSlug: {[key: string]: Position} = {

}

const locatedOutfitSlugs = Object.keys(positionsByCharacterSlug);

export function getRandomLocatedOutfitSlug(): string {
	return locatedOutfitSlugs[Math.floor(Math.random() * locatedOutfitSlugs.length)];
}
