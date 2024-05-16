export const Taxi = (distance: number, time: number) => {
    const cDistance: number = (Math.ceil(distance / 0.5) * 0.5) * 4
    const cTime: number = Math.ceil(time) * 1
    console.log(cDistance, cTime)
    return cDistance + cTime
}

export const calculateFare = (distance: number, waitingTime: number) => {
    return 4 * distance + waitingTime
}

export const roundDistance = (distance: number) => {
    return Math.ceil(distance * 2) / 2
}

export const roundWaitingTime = (waitingTime: number) => {
    return Math.ceil(waitingTime)
}

export const fare = (distance: number, waitingTime: number) => {
    return calculateFare(roundDistance(distance), roundWaitingTime(waitingTime))
}

export const minimumFare = (fare: number) => {
    return fare > 35 ? fare : 35
}
