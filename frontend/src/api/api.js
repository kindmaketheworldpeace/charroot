import {$axios} from './axios'

export default {
   entry: params => {
        return $axios.post('entry/', params)
    },

}
