import React from 'react';
import {useTranslation} from "react-i18next";

function Search(props) {

    const { t } = useTranslation();

    return (
        <div>
            <input className="text-lg rounded-full outline  placeholder:text-lg outline-2 outline-blue-50 bg-[#111011] bg-opacity-20 px-4 py-1 min-w-[250px]" placeholder={t('search')}/>
        </div>
    );
}

export default Search;