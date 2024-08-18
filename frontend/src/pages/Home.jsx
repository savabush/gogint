import React from 'react';
import Header from "../components/Header.jsx";
import Intro from "../components/Intro.jsx";
import Footer from "../components/Footer.jsx";
import Search from "../components/Search.jsx";
import Posts from "../components/Posts.jsx";
import {useTranslation} from "react-i18next";

function Home() {

    const { t } = useTranslation();

    return (
        <div className="font-JetBrainsMonoBold">
            <Header/>

            <Intro/>
            <hr className="mt-5 border-gray-700"></hr>

            <p className="justify-evenly flex text-3xl mt-5 flex-wrap gap-4">
                {t("popularPosts")}
                <Search/>
            </p>
            <div className="items-center justify-center flex">
                <Posts isPopular={true}/>
            </div>
            <Footer/>
        </div>
    );
}

export default Home;