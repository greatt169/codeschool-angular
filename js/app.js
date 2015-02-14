(function () {
    var app = angular.module('gemStore', []);
    var gems = [
        {
            name: 'Azurite',
            price: 2.95,
            canPurchase: false,
            soldOut: false
        },
        {
            name: 'Bloodstone',
            price: 2.95,
            canPurchase: true,
            soldOut: false
        },
        {
            name: 'Zircon',
            price: 5,
            canPurchase: true,
            soldOut: false
        }

    ];

    app.controller("StoreController", function () {
        this.products = gems;
    })

})();
