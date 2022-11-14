import {defineStore} from 'pinia'

const useAppStore = defineStore({
  id: 'app',
  state: () => ({
    token: '',
  }),
  getters: {
    getToken(): string {
      return this.token
    },
  },
  actions: {
    setToken(token: string) {
      this.token = token
    },
  },
})

export default useAppStore
