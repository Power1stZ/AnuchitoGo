export const Taxi = (distance:number,time:number) => {
    const cDistance:number = (Math.ceil(distance / 0.5) * 0.5)*4
    const cTime:number = Math.ceil(time)*1
    console.log(cDistance,cTime)
    return cDistance + cTime
}
