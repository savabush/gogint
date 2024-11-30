import React from 'react';
import ToMarkdown from '../ToMarkdown.jsx';

function Post(props) {

    const file = `
## Цель

- Получить опыт интеграции с [[Машинное обучение|ML]] моделями
- Получить деньги
- Опыт коммерческого стартапа
## Задачи

- Разработать чат бота на [[Микросервисы|микросервисной архитектуре]]
- Выложить продукт в мир, получить первых клиентов
- Интегрироваться с разными CRM и мессенджерами

## Стек

1. [[FastAPI]]
2. [[Kafka]]
3. [[Elasticsearch]]
4. [[Postgresql]]
5. [[Redis]]
6. [[BERT]]
7. [[ELK]]
8. [[Sentry]]
9. [[Graphana]]
## Архитектура #architecture

![[Pasted image 20241108201200.png]]

## Примечания

Пока пусто

    `

        return (
            <div className='flex items-center justify-center'>
                <ToMarkdown markdownText={file}/>
            </div>
        );
    }

export default Post;
