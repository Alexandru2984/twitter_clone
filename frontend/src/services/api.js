import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add token to requests if it exists
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export const authService = {
  register: (data) => api.post('/register', data),
  login: (data) => api.post('/login', data),
  getProfile: () => api.get('/profile'),
};

export const tweetService = {
  getTweets: () => api.get('/tweets'),
  getUserTweets: (username) => api.get(`/tweets/user/${username}`),
  createTweet: (content) => api.post('/tweets', { content }),
  likeTweet: (id) => api.post(`/tweets/${id}/like`),
  deleteTweet: (id) => api.delete(`/tweets/${id}`),
};

export const userService = {
  getUserProfile: (username) => api.get(`/users/${username}`),
  followUser: (username) => api.post(`/users/${username}/follow`),
};

export default api;
