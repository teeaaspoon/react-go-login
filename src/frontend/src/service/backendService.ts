import axios, {AxiosResponse} from "axios";

type BackendService = {
  login: (email: string, password: string) => Promise<AxiosResponse>;
  logout: () => Promise<AxiosResponse>;
  getDashboard: () => Promise<AxiosResponse>;
  getUser: () => Promise<AxiosResponse>;
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
  },

  getDashboard: () => {
    return instance.get("dashboard", {
      withCredentials: true,
      headers: {
        'Csrf-Token': localStorage.getItem("CSRFToken"),
      }
    })
  },

  getUser: () => {
    return instance.get("user", {
      withCredentials: true,
      headers: {
        'Csrf-Token': localStorage.getItem("CSRFToken"),
      }
    })
  },
}
