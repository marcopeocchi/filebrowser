import { Buffer } from "buffer"

export function getRemote() {
  return import.meta.env.DEV ? 'http://localhost:8080' : ''
}

export const encodeHexString = (value: string) =>
  Buffer.from(value).toString('hex')

export const decodeHexString = (value: string) =>
  Buffer.from(value, 'hex').toString()