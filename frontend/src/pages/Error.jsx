import React from 'react';
import {useTranslation} from "react-i18next";

function Error() {

    const { t } = useTranslation();

    return (
        <div className="flex flex-col items-center justify-center h-[80vh] mx-5">
            <h1 className="text-9xl font-bold">404</h1>
            <h2 className="text-5xl">{t("notFound")}</h2>
        </div>
    );
}

export default Error;