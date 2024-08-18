import React from 'react';
import {useTranslation} from "react-i18next";
import Links from "./Links.jsx";

function Footer(props) {

    const { t } = useTranslation();
    return (
        <div>
            <hr className="mt-5 border-blue-800"></hr>
            <div className="h-[100px] flex items-center justify-center">
                <ul className="flex gap-16 opacity-40 px-8">
                    <Links />
                </ul>
            </div>
        </div>
    );
}

export default Footer;