import './App.css'
import Home from "./pages/Home.jsx";
import Error from "./pages/Error.jsx";
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import Header from "./components/Header.jsx";
import React from "react";
import Footer from "./components/Footer.jsx";
import PostsPage from "./pages/PostsPage.jsx";
import PostPage from "./pages/PostPage.jsx";
import ArticlesPage from "./pages/ArticlesPage.jsx";
import ArticlePage from "./pages/ArticlePage.jsx";

const router = createBrowserRouter([
    {
        path: "/",
        element: <Home />,
    },
    {
        path: "/posts",
        element: <PostsPage />,
    },
    {
        path: "/posts/:id",
        element: <PostPage />,
    },
    {
        path: "/articles",
        element: <ArticlesPage />,
    },
    {
        path: "/articles/:id",
        element: <ArticlePage />,
    },
    {
        path: "*",
        element: <Error />,
    },
]);

function App() {
  return (
      <>
          <Header/>

          <RouterProvider router={router}/>
          <Footer />
      </>
  )
}

export default App
