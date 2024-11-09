import React from 'react';
import Intro from "../components/Intro.jsx";
import Search from "../components/Search.jsx";
import Posts from "../components/Post/Posts.jsx";
import {useTranslation} from "react-i18next";

function Home() {

    const { t } = useTranslation();

    return (
        <div>
            <Intro/>
            <hr className="mt-5 border-blue-800"></hr>

            <p className="justify-evenly flex text-4xl mt-5 flex-wrap gap-4 mx-8">
                {t("popularPosts")}
                <Search/>
            </p>
            <div className="items-center justify-center flex">
                <Posts isPopular={true}/>
            </div>
        </div>
    );
}

export default Home;