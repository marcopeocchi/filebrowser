export function getRemote() {
  return import.meta.env.DEV ? 'http://localhost:8080' : ''
}