import React from 'react';
import {useTranslation} from "react-i18next";
import Article from '../components/Article/Article.jsx';

function ArticlesPage() {

    const { t } = useTranslation();

    return (
        <div>
            <Article />
        </div>
    );
}

export default ArticlesPage;