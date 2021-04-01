import axios from 'axios'

export const getDataFromAPI = async () => {
    let data = await axios.get(`http://localhost:8081/feed`).then(response => response.data.response)
    return data
}
