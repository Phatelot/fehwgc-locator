const scaleFormatter = new Intl.NumberFormat('en-US', {maximumFractionDigits: 2})
const coordinateFormatter = new Intl.NumberFormat('en-US', {maximumFractionDigits: 1})

export function formatScale(scale: number) : string {
    return scaleFormatter.format(scale).replaceAll(',','')
}

export function formatCoordinate(coordinate: number) : string {
    return coordinateFormatter.format(coordinate).replaceAll(',','')
}
