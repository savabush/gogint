import React from 'react';
import Post from "../components/Post/Post.jsx";
import {useTranslation} from "react-i18next";
import { useParams } from 'react-router-dom';

function PostPage() {

    const { t } = useTranslation();

    const { id } = useParams();
    console.log(id)
    return (
        <div>
            <Post id={id}/>
        </div>
    );
}

export default PostPage;