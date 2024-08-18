import React from 'react';
import Intro from "../components/Intro.jsx";
import Search from "../components/Search.jsx";
import Posts from "../components/Posts.jsx";
import {useTranslation} from "react-i18next";
import Sort from "../components/Sort.jsx";

function PostsPage() {

    const { t } = useTranslation();

    return (
        <div>
            <div className="flex items-center flex-wrap justify-evenly mt-8 gap-4">
                <Sort />
                <Search />
            </div>
            <Posts isPopular={false}/>
        </div>
    );
}

export default PostsPage;