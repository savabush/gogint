import React from 'react';
import Search from "../components/Search.jsx";
import {useTranslation} from "react-i18next";
import Sort from "../components/Sort.jsx";
import Articles from "../components/Articles.jsx";

function ArticlesPage() {

    const { t } = useTranslation();

    return (
        <div>
            <div className="flex items-center flex-wrap justify-evenly mt-8 gap-4">
                <Sort />
                <Search />
            </div>
            <Articles />
        </div>
    );
}

export default ArticlesPage;