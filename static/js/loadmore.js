
;(function(w,$){
	
	var loadmore = { 
		/*单页加载更多 通用方法
		 * 
		 * @param callback 回调方法
		 * @param config 自定义参数
		 * */
		get : function(callback, config){
			var config = config ? config : {}; /*防止未传参数报错*/

			var counter = 0; /*计数器*/
			var pageStart = 0;
			var pageSize = config.size ? config.size : 10;

			/*默认通过点击加载更多*/
			$(document).on('click', '.js-load-more', function(){
				counter ++;
				pageStart = counter * pageSize;
				callback && callback(config, pageStart, pageSize);
			});
			
			/*通过自动监听滚动事件加载更多,可选支持*/
			config.isEnd = false; /*结束标志*/
			config.isAjax = false; /*防止滚动过快，服务端没来得及响应造成多次请求*/
			$(window).scroll(function(){
				
				/*是否开启滚动加载*/
				if(!config.scroll){
					return;
				}
				
				/*滚动加载时如果已经没有更多的数据了、正在发生请求时，不能继续进行*/
				if(config.isEnd == true || config.isAjax == true){
					return;
				}
				
				// /*当滚动到最底部以上100像素时， 加载新内容*/
				// if ($(document).height() - $(this).scrollTop() - $(this).height()<100){
				// 	counter ++;
				// 	pageStart = counter * pageSize;
				//
				// 	callback && callback(config, pageStart, pageSize);
				// }
			});

			/*第一次自动加载*/
			callback && callback(config, pageStart, pageSize);
		},
			
	}

	$.loadmore = loadmore;
})(window, window.jQuery || window.Zepto);

Date.prototype.Format = function (fmt) { //author: meizz
	var o = {
		"M+": this.getMonth() + 1, //月份
		"d+": this.getDate(), //日
		"h+": this.getHours(), //小时
		"m+": this.getMinutes(), //分
		"s+": this.getSeconds(), //秒
		"q+": Math.floor((this.getMonth() + 3) / 3), //季度
		"S": this.getMilliseconds() //毫秒
	};
	if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
	for (var k in o)
		if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
	return fmt;
}