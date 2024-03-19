### Использованные библиотеки и инструменты
- **самописная** библиотека для работы с PostgreSQL: https://github.com/timohahaa/postgres

Так же приложил sql-скрипты для таблиц в папке `migrations`

### Необязательно, но было бы круто
*Хорошо, если добавите поддержку конфигурации итоговой реализации. Параметры — на ваше усмотрение.*
Как я это реализовал:
- конфигурация хранилища данных через библиотеку для работы с Postgres (количество подключений, настройки подключения и прочее)
- конфигурация параметров N и K на этапе создания объекта FloodController (эти параметры в основное приложение можно передавать например через какой нибудь `config.yaml` файл)

# ХОД МЫСЛЕЙ
Изначально я хотел сделать "быструю и простую" реализацию с помощью пакета golang.org/x/time/rate, но потом быстро столкнулся с проблемой хранения данных. По условию нужно было предусмотреть хранилище данных (чтобы флуд-контроль запускать в разных приложениях), но я столкнулся с проблемой: как хранить специфичные для `time/rate` пакета данные и иметь доступ к ним из разных приложений? 
Данные легко можно бы было хранить в памяти с помощью мапы и мьютекса, но это ограничивает доступ к ним из разных приложений. Получилась бы слишком сложная структура флуд-контроля, я хотел найти что-то попроще.

Вторая мысль - можно "написать" пакет `time/rate` самому - только хранить данные не в памяти, а в стороннем хранилище. Я выбрал для этого PostgreSQL, потому что оно мне лучше всего знакомо. 
А дальше все просто - нужно просто сохранить ID пользователя и время прихода запроса от него в базу данных. Тогда при получении нового запроса достаточно посчитать по временному штампу в базе, сколько запросов было сделано этим пользователем за последние N секунд. Ну и после обработать результат. 
Есть у этого решения и минус - храним в базе данных информацию о куче разных запросов. Почему это не так плохо? Эти данные потом отдать для анализа аналитикам или просто с помощью cronjob переодически подчищать базу от старых запросов - чтобы она не стала слишком большой.

По поводу конфигураций - думаю, что возможности задавать значения N и K на этапе создания объекта флуд-контроля достаточно. Скорее всего он будет создаваться один на все приложение и использоваться где нибудь в middleware входящих запросов. Сами параметры такой вещи, как флуд-контроль, как мне показалось, меняются редко, но расширить итоговую реалзацию позволив менять находу значения N и K не так сложно.