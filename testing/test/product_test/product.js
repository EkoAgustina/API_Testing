const request = require('supertest');
const expect = require('chai').expect;
const env = require('../../env.json');
const productData = require('../../test-data/product/testdata_product.json');
const setup = require('../user_test/user')


const header = {
    'Accept': 'application/json',
    'Content-Type': 'application/json'
}

describe('PRODUCT', () =>{

    it('Successfully added product', () => {
        return new Promise((resolve,reject) => {
            setup.agent
            .post('/api/addProduct')
            .set(header)
            .send(productData)
            .expect(function (res) {
                if (res.statusCode == 200){
                    console.log('Adding products, responseCode:'+res.body.responseCode)
                }
                else{
                    reject(new Error('Adding products response code not expected, responseCode:'+res.body.responseCode))
                }

            })
            .end((err,ress)=>{
                if(err){
                    reject(new Error('There an error: '+res.error))
                }
                else{
                    resolve(ress)
                }
            })
        })
    })

    it('Successfully get product data', () => {
        return new Promise((resolve,reject) => {
            setup.agent
            .get('/api/products')
            .set(header)
            .expect(function (res) {
                if (res.statusCode == 200){
                    console.log('Get product, responseCode:'+res.body.responseCode)
                }
                else{
                    reject(new Error('get product response code not expected, responseCode:'+res.body.responseCode))
                }

            })
            .end((err,ress)=>{
                if(err){
                    reject(new Error('There an error: '+res.error))
                }
                else{
                    resolve(ress)
                }
            })
        })
    })


})