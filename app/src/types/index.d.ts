type DirectoryEntry = {
  name: string
  path: string
  size: number
  shaSum: string
  modTime: string
  isVideo: boolean
  upperLevel: string
  isDirectory: boolean
}

type APIResponse = {
  list: DirectoryEntry[]
  basePathLength: number
}

type DeleteRequest = Pick<DirectoryEntry, 'path' | 'shaSum'>

type PlayRequest = Pick<DirectoryEntry, 'path'>