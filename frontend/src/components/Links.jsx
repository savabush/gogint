import React from 'react';


function Links(props) {
    return (
        <>
            <li>
                <a href="https://github.com/savabush" target="_blank" rel="noreferrer" className="icon-link">
                    <img src="https://img.icons8.com/?size=100&id=12599&format=png&color=000000"
                         alt="Github Icon" className="icon w-12"/>
                </a>
            </li>
            <li>
                <a href="https://t.me/sava_dev" target="_blank" rel="noreferrer" className="icon-link">
                    <img src="https://img.icons8.com/fluency-systems-filled/48/000000/telegram-app.png"
                         alt="Telegram Icon" className="icon w-12"/>
                </a>
            </li>
            <li>
                <a href="https://www.instagram.com/sssava.b/" target="_blank" rel="noreferrer"
                   className="icon-link">
                    <img src="https://img.icons8.com/fluency-systems-filled/96/000000/instagram-new.png"
                         alt="Instagram Icon" className="icon w-12"/>
                </a>
            </li>
            <li>
                <a href="mailto:vatka1337@gmail.com" target="_blank" rel="noreferrer" className="icon-link">
                    <img src="https://img.icons8.com/fluency-systems-filled/96/000000/email.png" alt="Email Icon"
                         className="icon w-12"/>
                </a>
            </li>
            <style jsx>{`
                        .icon-link {
                            filter: invert(100%);
                        }
                    `}</style>
            </>
    );
}

export default Links;