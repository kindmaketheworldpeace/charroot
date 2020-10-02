import api from '@/api/api'

const state = {
}

const getters = {
}

const mutations = {
}

const actions = {
  entry({commit, state}, param) {
        return api.entry(param)
    },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
