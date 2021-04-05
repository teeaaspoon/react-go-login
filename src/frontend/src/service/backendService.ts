import axios, {AxiosResponse} from "axios";

type BackendService = {
  login: (email: string, password: string) => Promise<AxiosResponse>;
  logout: () => Promise<AxiosResponse>;
}

const instance = axios.create({
  baseURL: "http://localhost:8000/api/"
})

export const backendService: BackendService = {
  login: (email: string, password: string) => {
    return instance.post("login", {email, password}, {
      withCredentials: true,
      headers: {
        'Content-Type': 'application/json'
      }
    });
  },

  logout: () => {
    return instance.post("logout", {}, {
      withCredentials: true,
      headers: {
        'Content-Type': 'application/json'
      }
    })
  }
}
