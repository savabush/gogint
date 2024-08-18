import React from 'react';

function Articles(props) {

    const data = {
        "result": [
            {
                id: "2",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя",
                summary: "Краткое содержание из ChatGPT (условно)Краткое содержание из ChatGPT (условно) 4Кре содераткое содержание из ChatGPT (условно) 4"
            },
            {
                id: "3",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4"
            },
            {
                id: "4",
                img: "https://avatars.mds.yandex.net/i?id=3de58bd676e0e279fca5c9e68a825f64_l-9246913-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4"
            },
            {
                id: "5",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPTсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (усл (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4"
            },
            {
                id: "6",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPTсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (усл (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4"
            },
            {
                id: "7",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPTсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (усл (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4"
            },
            {
                id: "8",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPTсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (усл (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4"
            },
            {
                id: "9",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPTсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (усл (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4"
            },
            {
                id: "10",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPTсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (усл (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4"
            },
            {
                id: "11",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPTсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (усл (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4"
            },
        ]
    }


    return (
        <div>
            <ul className="flex flex-wrap justify-center">
                {data.result.map((item) => (
                    <div key={item.id} className="items-center flex flex-col p-8">
                        <a href="#"><img
                            className="w-[350px] rounded-2xl mb-2" src={item.img} alt={item.name.toLowerCase()}/></a>
                            <p><a href="#" className="hover:font-JetBrainsMonoExtraBold text-white text-3xl">{item.name.toLowerCase()}</a></p>
                        <span className="max-w-[350px] max-h-[200px] line-clamp-6 text-lg">{item.summary.toLowerCase()}</span>
                    </div>
                ))}
            </ul>
        </div>
    );
}

export default Articles;