curl -X POST 'http://localhost:8123/ordens' -d '{
    "cliente": "José do Teste",
    "produtos": [
        {
            "nome": "Produto 1",
            "quantidade": 2,
            "cores": [
                "Azul",
                "Vermelho"
            ],
            "valor": 10.00
        },
        {
            "nome": "Produto 2",
            "quantidade": 1,
            "valor": 20.00
        }
    ]
}'

curl 'http://localhost:8123/produtos'