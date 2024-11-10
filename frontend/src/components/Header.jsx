import React from 'react';
import {supportedLngs} from "../i18n/Config.jsx";
import {useTranslation} from "react-i18next";
import Links from "./Links.jsx";

function Header(props) {

    const { t, i18n } = useTranslation();

    return (
        <div>
            <div className='flex items-center justify-between px-16 pt-4 flex-wrap'>
                <h2 className='text-5xl w-[280px] mb-4'>
                    <a href="/" className="text-white hover:font-JetBrainsMonoExtraBold">{'<'}{t('main')}{'>'}</a>
                </h2>

                <ul className='flex items-center text-5xl flex-wrap mb-4'>
                    <li className="mb-1 mr-8"><a href="/posts" className="text-white hover:font-JetBrainsMonoExtraBold">{t('posts')}</a></li>
                    <li className="mb-1 ml-8"><a href="/articles" className="text-white hover:font-JetBrainsMonoExtraBold">{t('articles')}</a>
                    </li>
                </ul>

                <ul className='flex gap-4 mb-4 items-center'>
                    <select
                        className="hover:outline-blue-900 rounded-2xl px-3 mr-3 h-8 border-r-8 border-transparent pr-1 text-lg outline outline-neutral-700 bg-[#111011] appearance-none"
                        value={i18n.resolvedLanguage}
                        onChange={(e) => i18n.changeLanguage(e.target.value)}
                    >
                        {Object.entries(supportedLngs).map(([code, name]) => (
                            <option value={code} key={code}>
                                {name}
                            </option>
                        ))}
                    </select>
                    <Links />
                </ul>

            </div>
            <hr className="mt-5 border-blue-800"></hr>
        </div>

);
}

export default Header;