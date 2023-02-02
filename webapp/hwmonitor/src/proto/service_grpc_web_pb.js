/**
 * @fileoverview gRPC-Web generated client stub for main
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.20.3
// source: proto/service.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.main = require('./service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.main.HardwareMonitorClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.main.HardwareMonitorPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.main.EmptyRequest,
 *   !proto.main.HardwareStats>}
 */
const methodDescriptor_HardwareMonitor_Monitor = new grpc.web.MethodDescriptor(
  '/main.HardwareMonitor/Monitor',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.main.EmptyRequest,
  proto.main.HardwareStats,
  /**
   * @param {!proto.main.EmptyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.main.HardwareStats.deserializeBinary
);


/**
 * @param {!proto.main.EmptyRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.main.HardwareStats>}
 *     The XHR Node Readable Stream
 */
proto.main.HardwareMonitorClient.prototype.monitor =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/main.HardwareMonitor/Monitor',
      request,
      metadata || {},
      methodDescriptor_HardwareMonitor_Monitor);
};


/**
 * @param {!proto.main.EmptyRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.main.HardwareStats>}
 *     The XHR Node Readable Stream
 */
proto.main.HardwareMonitorPromiseClient.prototype.monitor =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/main.HardwareMonitor/Monitor',
      request,
      metadata || {},
      methodDescriptor_HardwareMonitor_Monitor);
};


module.exports = proto.main;
