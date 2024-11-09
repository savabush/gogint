import React from 'react';
import Markdown from 'react-markdown';
import remarkGfm from 'remark-gfm'
import style from './markdown-styles.module.css'

function ToMarkdown(props) {

    
    const file = props.markdownText

        return (
            <div className='flex items-center justify-center'>
                <div className='text-2xl mt-10 mb-10'>
                    <Markdown className={style.reactMarkDown} remarkPlugins={[remarkGfm]}>{file}</Markdown>
                </div>
            </div>
        );
    }

export default ToMarkdown;
