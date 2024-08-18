import React, {useEffect, useState} from 'react';
import GlobalApi from "../services/GlobalApi.jsx";
import {useTranslation} from "react-i18next";

function Intro(props) {

    const { t } = useTranslation();

    useEffect(() => {
        getAboutMe();
    }, [])

    const getAboutMe = () => {
        GlobalApi.getAboutMe.then(resp => {
            console.log(resp)
            // const data = resp.data;

        });
    }
    const data = {
        name: 'Saveliy Bushkov',
        about: 'i\'ve already developed a bunch of commercial projects, mostly i\'m running around the backend, but i\'m learning a little bit on the full stack',
        profession: 'backend developer',
    }

    return (
        <div>
            <ul className="flex justify-evenly mt-10 items-center flex-wrap">
                <img className="rounded-full w-80 mb-5 mx-8" src="https://avatars.githubusercontent.com/u/87959063?v=4"
                     alt="My Picture"/>
                <ul className="mx-[5%] flex flex-col items-center">
                    <li>
                        <ul className="flex-col">
                            <li><p className="text-5xl">{data.name}</p></li>
                            <li><span
                                className="mt-4 opacity-40 text-2xl">{data.profession}</span>
                            </li>
                            <li className="break-words text-wrap"><span
                                className="mt-4 text-2xl">{data.about}</span>
                            </li>
                        </ul>
                    </li>

                    <li className="justify-center">
                        <ul className="flex gap-8">
                            <li><a href="https://t.me/sava_dev" className="text-white hover:font-JetBrainsMonoExtraBold hover:text-white no-underline"><button className="text-2xl rounded-full px-4 py-1 mt-12">{t("contact me")}</button></a></li>
                            <li><button className="text-2xl rounded-full px-4 py-1 mt-12 hover:font-JetBrainsMonoExtraBold">{t("curriculum vitae")}</button></li>
                        </ul>
                    </li>
                </ul>
            </ul>
        </div>
    );
}

export default Intro;