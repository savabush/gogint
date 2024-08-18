import React from 'react';

function Posts(props) {

    const data = {
        "result": [
            {
                id: "2",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя",
                summary: "Краткое содержание из ChatGPT (условно)Краткое содержание из ChatGPT (условно) 4Кре содераткое содержание из ChatGPT (условно) 4"
            },
            {
                id: "1",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4"
            },
            {
                id: "1",
                img: "https://avatars.mds.yandex.net/i?id=89564501da7fed05c2040432e9175e33_l-8972573-images-thumbs&n=13",
                name: "Тестовое имя 2",
                summary: "Краткое содержание из ChatGPT (условно) 4 Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4Краткое содержание из ChatGPT (условно) 4"
            },
            {
                id: "1",
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
                            <p><a href="#" className="hover:font-JetBrainsMonoExtraBold text-white">{item.name.toLowerCase()}</a></p>
                        <span className="max-w-[350px] max-h-[200px] line-clamp-5">{item.summary.toLowerCase()}</span>
                    </div>
                ))}
            </ul>
        </div>
    );
}

export default Posts;