import axios from "axios";


const BASE_URL = 'http://localhost:8000/api/v1/';

const getAboutMe = axios.get(`${BASE_URL}aboutMe`);
const getIntro = axios.get(`${BASE_URL}intro`);
const getPost = axios.get(`${BASE_URL}posts/`);
const getPosts = axios.get(`${BASE_URL}posts`);
const getArticles = axios.get(`${BASE_URL}articles`);

export default {
    getAboutMe,
    getIntro,
    getPost,
    getPosts,
    getArticles
}