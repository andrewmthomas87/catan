export interface ILogger {
	trace(context: string, ...args: any[]): void
	info(context: string, ...args: any[]): void
	debug(context: string, ...args: any[]): void
	warn(context: string, ...args: any[]): void
	error(context: string, ...args: any[]): void
}

export interface ILoggerFactory {
	getLogger(name: string): ILogger
}

export class LoggerFactory implements ILoggerFactory {
	public getLogger(name: string): ILogger {
		return {
			trace: this._trace.bind(null, name),
			debug: this._debug.bind(null, name),
			info: this._info.bind(null, name),
			warn: this._warn.bind(null, name),
			error: this._error.bind(null, name),
		}
	}

	private _trace(name: string, context: string, ...args: any[]) {
		console.trace(`%c[T] ${name}:${context}:`, 'color: #FFC107;', ...args)
	}

	private _debug(name: string, context: string, ...args: any[]) {
		console.debug(`%c[D] ${name}:${context}:`, 'color: #E91E63;', ...args)
	}

	private _info(name: string, context: string, ...args: any[]) {
		console.info(`%c[I] ${name}:${context}:`, 'color: #673AB7;', ...args)
	}

	private _warn(name: string, context: string, ...args: any[]) {
		console.warn(`%c[W] ${name}:${context}:`, 'color: #FF5722;', ...args)
	}

	private _error(name: string, context: string, ...args: any[]) {
		console.error(`%c[E] ${name}:${context}:`, 'color: #F44336;', ...args)
	}
}
