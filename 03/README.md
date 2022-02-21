Всем хотя бы раз в жизни приходилось перекладывать JSON. Вот и для нового проекта под названием "Единое хранилище" необходимо переложить магазинные фиды. Для размещения на Яндекс.Маркете магазины передают товары из своего ассортимент в JSON-файлах. Одно товарное предложение описывается так:

{
    "offer_id": <string>,
    "market_sku": <int>,
    "price": <int>
}
где offer_id - уникальный среди всех фидов магазина идентификатор предложения, market_sku - идентификатор товара на Яндекс.Маркете, price - стоимость товара.

Весь фид магазина представляет собой JSON и выглядит так:

{
    "offers": [<offer>, <offer>, ...]
}
Вас попросили написать программу, которая объединит все фиды одного магазина в единый фид и отсортирует товары в порядке неубывания их стоимостей, а при их равенстве - по offer_id.

Формат вывода
Выходной поток должен содержать один JSON-документ, удовлетворяющий формату товарного фида. Количество строк в выходном файле и табуляция не имеют значения.

Пример
Ввод
2
{"offers": [{"offer_id": "offer1", "market_sku": 10846332, "price": 1490}, {"offer_id": "offer2", "market_sku": 682644, "price": 499}]}
{"offers": [{"offer_id": "offer3", "market_sku": 832784, "price": 14000}]}
Вывод
{"offers":[{"market_sku":682644,"offer_id":"offer2","price":499},{"market_sku":10846332,"offer_id":"offer1","price":1490},{"market_sku":832784,"offer_id":"offer3","price":14000}]}

#include "json.hpp"
Для решений на golang доступны все стандартные пакеты, включая encoding/json, sort и другие.