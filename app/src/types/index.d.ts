type DirectoryEntry = {
  name: string
  path: string
  size: number
  shaSum: string
  modTime: string
  isVideo: boolean
  isDirectory: boolean
}

type APIResponse = {
  list: DirectoryEntry[]
  upperLevelPath: string
}

type DeleteRequest = Pick<DirectoryEntry, 'path' | 'shaSum'>

type PlayRequest = Pick<DirectoryEntry, 'path'>