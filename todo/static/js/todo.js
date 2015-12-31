/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/

function TaskCtrl($scope, $http) {
  $scope.tasks = [];
  $scope.working = false;

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
    $scope.working = false;
  };

  var refresh = function() {
    return $http.get('/task/').
      success(function(data) { $scope.tasks = data.Tasks; }).
      error(logError);
  };

  $scope.addTodo = function() {
    $scope.working = true;
    $http.post('/task/', {Title: $scope.todoText}).
      error(logError).
      success(function() {
        refresh().then(function() {
          $scope.working = false;
          $scope.todoText = '';
        })
      });
  };

  $scope.toggleDone = function(task) {
    data = {ID: task.ID, Title: task.Title, Done: !task.Done}
    $http.put('/task/'+task.ID, data).
      error(logError).
      success(function() { task.Done = !task.Done });
  };

  refresh().then(function() { $scope.working = false; });
}

function PostCtrl($scope, $http) {
  $scope.posts = [];

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
  };

  var refresh = function() {
    return $http.get('/post/').
      success(function(data) { $scope.posts = data; }).
      error(logError);
  };

  $scope.addPost = function() {
    $http.post('/post/', {Title: $scope.postText}).
      error(logError).
      success(function() {
      });
  };

  $scope.delPost = function() {
      $http.delete('/post/'+$scope.postText).
        error(logError).
        success(function() {
        });
    };

    $scope.updatePost = function() {
        $http.put('/post/'+$scope.postText, {Body: "hahaha"}).
          error(logError).
          success(function() {
          });
      };

  refresh().then(function() { });
}