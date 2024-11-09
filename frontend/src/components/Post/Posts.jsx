import React from 'react';

function Posts(props) {

    const isPopular = props.isPopular;

    const data = {
        "result": [
            {
                id: "2",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя",
                summary: "Краткое содержание из ChatGPT (условно)Краткое содержание из ChatGPT (условно) 4Кре содераткое содержание из ChatGPT (условно) 4",
                updated_date: "2024-10-05 12:00:00"
            },
            {
                id: "3",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4",
                updated_date: "2024-10-05 12:00:00"
            },
            {
                id: "4",
                img: "https://avatars.mds.yandex.net/i?id=3de58bd676e0e279fca5c9e68a825f64_l-9246913-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4",
                updated_date: "2024-10-05 12:00:00"
            },
            {
                id: "5",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPTсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (усл (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4",
                updated_date: "2024-10-05 12:00:00"
            },
            {
                id: "6",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPTсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (усл (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4",
                updated_date: "2024-10-05 12:00:00"
            },
            {
                id: "7",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPTсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (усл (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4",
                updated_date: "2024-10-05 12:00:00"
            },
            {
                id: "8",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPTсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (усл (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4",
                updated_date: "2024-10-05 12:00:00"
            },
            {
                id: "9",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPTсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (усл (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4",
                updated_date: "2024-10-05 12:00:00"
            },
            {
                id: "10",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPTсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (усл (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4",
                updated_date: "2024-10-05 12:00:00"
            },
            {
                id: "11",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPTсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (услсловно) 4 Краткое содержание из ChatGPT (усл (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4",
                updated_date: "2024-10-05 12:00:00"
            },
        ]
    }

    const results = isPopular ? data.result.slice(0, 4) : data.result;
    console.log(window.location)
    if (isPopular) {
        return (
            <div>
                <ul className="flex flex-wrap justify-center">
                    {results.map((item) => (
                        <div key={item.id} className="items-center flex flex-col p-8">
                            <a href={window.location.href + "posts/" + item.id}><img
                                className="w-[350px] rounded-2xl mb-2" src={item.img}
                                alt={item.name.toLowerCase()}/></a>
                            <p><a href={window.location.href + "posts/" + item.id}
                                  className="hover:font-JetBrainsMonoExtraBold text-white text-3xl">{item.name.toLowerCase()}</a>
                            </p>
                            <span
                                className="max-w-[350px] max-h-[200px] line-clamp-6 text-lg">{item.summary.toLowerCase()}</span>
                        </div>
                    ))}
                </ul>
            </div>
        );
    } else {
        return (
            <div>
                <ul className="flex justify-center flex-col items-center">
                    {results.map((item) => (
                        <div key={item.id} className="items-end flex flex-col p-8">
                            <a href={window.location.pathname + "/" + item.id}><img
                                className="w-[100%] max-w-[700px] h-auto rounded-2xl mb-2" src={item.img}
                                alt={item.name.toLowerCase()}/></a>
                            <p className='self-center'><a href={window.location.pathname + "/" + item.id}
                                  className="hover:font-JetBrainsMonoExtraBold text-white text-3xl">{item.name.toLowerCase()}</a>
                            </p>
                            <span
                                className="max-w-[700px] max-h-[200px] line-clamp-6 text-lg">{item.summary.toLowerCase()}</span>
                            <span
                                className="text-lg mt-6 text-gray-600">{item.updated_date}</span>
                        </div>
                    ))}
                </ul>
            </div>
        );
    }

}

export default Posts;