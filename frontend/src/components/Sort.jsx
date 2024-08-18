import React, { useState } from 'react';
import {useTranslation} from "react-i18next";

const Sort = (props) => {

    const { t } = useTranslation();

    const [sortState, setSortState] = useState({
        sortBy: 'created',
        sortOrder: 'desc',
    });

    const handleSort = (event) => {
        const sortBy = event.target.value;
        const sortOrder = event.target.options[event.target.selectedIndex].dataset.order;
        setSortState({ sortBy, sortOrder });
        // props.onSort(sortBy, sortOrder);
    };

    return (
        <div className="flex">
            <select id="sort" className="rounded-full px-4 outline-2 outline-neutral-600 text-neutral-400 outline bg-[#111011] bg-opacity-20 appearance-none border-transparent text-lg" value={sortState.sortBy} onChange={handleSort}>
                <option className="bg-[#111011]" value="created-desc" data-order="desc">{t("created desc")}</option>
                <option className="bg-[#111011]" value="created-asc" data-order="asc">{t("created asc")}</option>
                <option className="bg-[#111011]" value="title-desc" data-order="desc">{t("title desc")}</option>
                <option className="bg-[#111011]" value="title-asc" data-order="asc">{t("title asc")}</option>
            </select>
        </div>
    );
};

export default Sort;