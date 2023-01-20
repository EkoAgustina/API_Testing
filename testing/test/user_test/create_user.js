const request = require('supertest');
const expect = require('chai').expect;
const env = require('./../../env.json');
const userData = require('./../../test-data/user/testdata_user.json');

const header = {
    'Accept': 'application/json',
    'Content-Type': 'application/json'
}

describe('USER', () =>{

    it('Successfully created user', () => {
        return new Promise((resolve,reject) => {
            request(env.baseUrl)
            .post('/register')
            .set(header)
            .send(userData)
            .expect(function (res) {
                if (res.statusCode != 200){
                    reject(new Error('Response Code not Expected, responseCode: '+res.body.responseCode))
                }
                else{
                    console.log('Successfully, responseCode: '+res.body.responseCode)
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
    });


})