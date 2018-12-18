// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

		
	$scope.carHistory = function(){
		var id = $scope.sn;
		console.log(id)
		appFactory.carHistory(id, function(data){
		
			$scope.car_History = data;			
		});
	}
	
	$scope.Newcar = function(){

		appFactory.Newcar($scope.record_car, function(data){
			$scope.record_car = data;
		
		});
	}

	$scope.Change = function(){

		appFactory.Change($scope.record_change, function(data){
			$scope.record_change = data;
		
		});
	}

	$scope.Repair = function(){

		appFactory.Repair($scope.record_repair, function(data){
			$scope.record_repair = data;
			
		});
	}

	$scope.Accident = function(){

		appFactory.Accident($scope.record_accident, function(data){
			$scope.record_accident = data;
			
		});
	}


});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

	factory.carHistory = function(id, callback){
    	$http.get('/history/'+id).success(function(output){
			callback(output)
		});
	}

  	factory.newcar = function(data, callback){

		var record_car = data.id + "-" + data.manufacture + "-" + data.factory + "-" + data.name + "-" + data.vehicletype + "-" + data.volume + "-" + data.fuel;

    	$http.get('/init_car/'+record_car).success(function(output){
			callback(output)
		});
	}

	factory.change = function(data, callback){

		var record_change = data.id + "-" + data.government + "-" + data.plate + "-" + data.owner + "-" + data.tradehistory + "-" + data.price;

    	$http.get('/change_car/'+record_change).success(function(output){
			callback(output)
		});
	}

	factory.repair = function(data, callback){

		var record_repair = data.id + "-" + data.repair + "-" + data.repaircosts + "-" + data.shop;

    	$http.get('/repair_car/'+record_repair).success(function(output){
			callback(output)
		});
	}

	factory.accident = function(data, callback){

		var record_accident = data.id + "-" + data.accident  + "-" + data.handdlingcosts+ "-" + data.insurance;

    	$http.get('/accident_car/'+record_accident).success(function(output){
			callback(output)
		});
	}


	return factory;
});





