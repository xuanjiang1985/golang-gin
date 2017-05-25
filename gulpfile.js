 var gulp = require('gulp'),  
    runSequence = require('run-sequence'),  
    rev = require('gulp-rev'),  
    revCollector = require('gulp-rev-collector'),
    uglify = require('gulp-uglify'),
    minify = require('gulp-clean-css'),
    del = require('del');
//定义css、js源文件路径  
var cssSrc = './assets/css/*.css',  
    jsSrc = './assets/js/*.js';  

//delete hash files
gulp.task('delete', function(){
    del([
            './public/build/css/*.css',
            './public/build/js/*.js'
        ]);
}); 
  
//CSS生成文件hash编码并生成 rev-manifest.json文件名对照映射  
gulp.task('revCss', function(){  
    return gulp.src(cssSrc)  
        .pipe(rev()) 
        .pipe(minify()) 
        .pipe(gulp.dest('./public/build/css'))  
        .pipe(rev.manifest())  
        .pipe(gulp.dest('./public/build/css'));  
});  
  
  
//js生成文件hash编码并生成 rev-manifest.json文件名对照映射  
gulp.task('revJs', function(){  
    return gulp.src(jsSrc)  
        .pipe(rev())
        .pipe(uglify())                               //给文件添加hash编码  
        .pipe(gulp.dest('./public/build/js'))  
        .pipe(rev.manifest())                       //生成rev-mainfest.json文件作为记录  
        .pipe(gulp.dest('./public/build/js'));  
});  
  
  
//Html替换css、js文件版本  
gulp.task('revHtmlCss', function () {  
    return gulp.src(['./public/build/css/*.json', './templates/base.html'])  
        .pipe(revCollector({
            replaceReved:true
        }))                         //替换html中对应的记录  
        .pipe(gulp.dest('./templates'));                     //输出到该文件夹中  
});  
gulp.task('revHtmlJs', function () {  
    return gulp.src(['./public/build/js/*.json', './templates/base.html'])  
        .pipe(revCollector({
            replaceReved:true
        }))  
        .pipe(gulp.dest('./templates'));  
});  
 
  
//开发构建  
gulp.task('default', function (done) {  
    condition = false;  
    //依次顺序执行  
    runSequence(
        ['delete'],
        ['revCss'],
        ['revHtmlCss'],
        ['revJs'],
        ['revHtmlJs'],
        done);  
});