import {Taxi} from './taxi'

describe('Taxi', () => {
  it('(int)KM and (int)TIME', () => {
    const result = Taxi(10,10)
    expect(result).toEqual(50)
  })

  it('(int)KM and (float)TIME > x.1',()=>{
    const result = Taxi(10,10.5)
    expect(result).toEqual(51)
  })

  it('x.5 > (float)KM > x.1 and (int)TIME',()=>{
    const result = Taxi(10.2,10)
    expect(result).toEqual(52)
  })

  it('x.5 < (float)KM < x.1 and (int)TIME',()=>{
    const result = Taxi(10.8,10)
    expect(result).toEqual(54)
  })

  it('x.5 == (float)KM and (int)TIME',()=>{
    const result = Taxi(10.5,10)
    expect(result).toEqual(52)
  })


})
